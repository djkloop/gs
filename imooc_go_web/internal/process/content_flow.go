package process

import (
	"encoding/json"
	"fmt"
	flow "github.com/s8sg/goflow/flow/v1"
	goflow "github.com/s8sg/goflow/v1"
	"gorm.io/gorm"
	"imooc_go_web/internal/dao"
	"time"
)

func ExecContentFlow(db *gorm.DB) {
	contentFlow := &contentFlow{
		contentDao: dao.NewContentDao(db),
	}
	fs := goflow.FlowService{
		Port:              8999,
		RedisURL:          "43.143.243.166:6379",
		RedisPassword:     "bS9@xG2?",
		WorkerConcurrency: 5,
	}
	_ = fs.Register("content-flow", contentFlow.flowHandle)
	err := fs.Start()
	if err != nil {
		panic(err)
	}
}

func ExecContentWork(db *gorm.DB) {
	contentFlow := &contentFlow{
		contentDao: dao.NewContentDao(db),
	}
	fs := goflow.FlowService{
		Port:              7788,
		RedisURL:          "localhost:6379",
		WorkerConcurrency: 4,
	}
	_ = fs.Register("content-flow", contentFlow.flowHandle)
	err := fs.StartWorker()
	if err != nil {
		panic(err)
	}
}

func ExecServer(db *gorm.DB) {
	fs := goflow.FlowService{
		Port:              7788,
		RedisURL:          "localhost:6379",
		WorkerConcurrency: 5,
	}
	contentFlow := &contentFlow{
		contentDao: dao.NewContentDao(db),
	}
	_ = fs.Register("content-flow", contentFlow.flowHandle)
	err := fs.StartServer()
	if err != nil {
		panic(err)
	}
}

type contentFlow struct {
	contentDao *dao.ContentDao
}

func (c *contentFlow) flowHandle(workflow *flow.Workflow, context *flow.Context) error {
	dag := workflow.Dag()
	// 创建三个节点
	dag.Node("input", c.input)
	dag.Node("verify", c.verify)
	dag.Node("finish", c.finish)
	// 创建分支
	branches := dag.ConditionalBranch(
		"branches",
		[]string{"category", "thumbnail", "pass", "format", "fail"},
		func(bytes []byte) []string {
			var data map[string]interface{}
			if err := json.Unmarshal(bytes, &data); err != nil {
				return nil
			}
			if data["approval_status"].(float64) == 2 {
				return []string{"category", "thumbnail", "pass", "format"}
			}
			return []string{"fail"}
		},
		flow.Aggregator(
			func(m map[string][]byte) ([]byte, error) {
				fmt.Println(m)
				return []byte("ok"), nil
			},
		),
	)
	branches["category"].Node("category", c.category)
	branches["thumbnail"].Node("thumbnail", c.thumbnail)
	branches["pass"].Node("category", c.pass)
	branches["format"].Node("format", c.format)
	branches["fail"].Node("fail", c.fail)

	dag.Edge("input", "verify")
	dag.Edge("verify", "branches")
	dag.Edge("branches", "finish")
	return nil
}

var startTime = time.Now()

// 加工流第一步，输入，查询数据库，将视频标题, 视频URL 传递给下一个节点
func (c *contentFlow) input(data []byte, option map[string][]string) ([]byte, error) {
	startTime = time.Now()
	fmt.Println("exec input")
	var input map[string]int
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	id := input["content_id"]
	detail, err := c.contentDao.First(id)
	if err != nil {
		return nil, err
	}
	result, err := json.Marshal(map[string]interface{}{
		"title":      detail.Title,
		"video_url":  detail.VideoURL,
		"content_id": detail.ID,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 加工流第二步，审核，根据视频标题和视频URL，判断视频是否符合要求，将审核结果传递给下一个节点
func (c *contentFlow) verify(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("exec verify")
	var detail map[string]interface{}
	if err := json.Unmarshal(data, &detail); err != nil {
		return nil, err
	}
	var (
		title    = detail["title"]
		videoURL = detail["video_url"]
		id       = detail["content_id"]
	)
	// 机审，人审
	if int(id.(float64))%2 == 0 {
		detail["approval_status"] = 3
	} else {
		detail["approval_status"] = 2
	}
	//detail["approval_status"] = 2
	fmt.Println(id, title, videoURL)
	return json.Marshal(detail)
}

// 加工流第三步，分类，根据视频标题和视频URL，判断视频分类，将分类结果传递给下一个节点
func (c *contentFlow) category(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("exec category")
	var input map[string]interface{}
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	contentID := int(input["content_id"].(float64))
	err := c.contentDao.UpdateByID(contentID, "category", "category")
	if err != nil {
		return nil, err
	}
	return []byte("category"), nil
}

// 加工流第四步，缩略图，根据视频标题和视频URL，判断视频缩略图，将缩略图结果传递给下一个节点
func (c *contentFlow) thumbnail(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("exec thumbnail")
	var input map[string]interface{}
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	contentID := int(input["content_id"].(float64))
	err := c.contentDao.UpdateByID(contentID, "thumbnail", "thumbnail")
	if err != nil {
		return nil, err
	}
	return []byte("thumbnail"), nil
}

// 加工流第五步，格式，根据视频标题和视频URL，判断视频格式，将格式结果传递给下一个节点
func (c *contentFlow) format(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("exec format")
	var input map[string]interface{}
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	contentID := int(input["content_id"].(float64))
	err := c.contentDao.UpdateByID(contentID, "format", "format")
	if err != nil {
		return nil, err
	}
	return []byte("format"), nil
}

// 加工流第六步，通过，根据视频标题和视频URL，判断视频通过，将通过结果传递给下一个节点
func (c *contentFlow) pass(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("exec pass")
	var input map[string]interface{}
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	contentID := int(input["content_id"].(float64))
	err := c.contentDao.UpdateByID(contentID, "approval_status", 2)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 加工流第七步，失败，根据视频标题和视频URL，判断视频失败，将失败结果传递给下一个节点
func (c *contentFlow) fail(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("exec fail")
	var input map[string]interface{}
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	contentID := int(input["content_id"].(float64))
	err := c.contentDao.UpdateByID(contentID, "approval_status", 3)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 加工流第八步，结束，根据视频标题和视频URL，判断视频结束
func (c *contentFlow) finish(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("exec finish")
	fmt.Println("result :", string(data))
	tc := time.Since(startTime) //计算耗时
	fmt.Printf("time cost = %v\n", tc)
	return data, nil
}
