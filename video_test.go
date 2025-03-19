package bilibili

import (
	"testing"
)

func TestGetVideoInfo(t *testing.T) {
	// 创建一个模拟的 Client 对象
	client := New()

	// 定义测试参数
	param := VideoParam{
		Bvid: "BV1B7X5YGEZJ", // 使用一个已知的 bvid 进行测试
	}

	// 调用 GetVideoInfo 方法
	videoInfo, err := client.GetVideoInfo(param)
	if err != nil {
		t.Fatalf("GetVideoInfo failed: %v", err)
	}

	// 检查返回的 VideoInfo 对象是否为空
	if videoInfo == nil {
		t.Error("Expected non-nil VideoInfo, got nil")
	}

	// 检查 VideoInfo 中的一些字段是否有效
	if videoInfo.Bvid != param.Bvid {
		t.Errorf("Expected Bvid to be %s, got %s", param.Bvid, videoInfo.Bvid)
	}

	if videoInfo.Title == "" {
		t.Error("Expected non-empty Title, got empty")
	}

	if videoInfo.Owner.Mid == 0 {
		t.Error("Expected non-zero Owner.Mid, got zero")
	}

	if videoInfo.Owner.Name == "" {
		t.Error("Expected non-empty Owner.Name, got empty")
	}
}
