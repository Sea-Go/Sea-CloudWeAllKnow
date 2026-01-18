package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"sea/config"
	"sea/zlog"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"go.uber.org/zap"
)

var cfg = config.Cfg.Ali

var client = openai.NewClient(
	option.WithAPIKey(cfg.APIKey),
	option.WithBaseURL(cfg.BaseURL),
)

func EmbeddingTxt(txt string) (*openai.CreateEmbeddingResponse, error) {
	res, err := client.Embeddings.New(context.TODO(), openai.EmbeddingNewParams{
		Input: openai.EmbeddingNewParamsInputUnion{
			OfString: openai.String(txt),
		},
		Model:          "text-embedding-v4",
		Dimensions:     openai.Int(2048),
		EncodingFormat: openai.EmbeddingNewParamsEncodingFormatFloat,
		User:           openai.String("user-neo"),
	})
	if err != nil {
		zlog.L().Error("embedding service fail", zap.Error(err))
	}
	return res, nil
}

func EmbeddingGraph(ty string, url string) (*openai.CreateEmbeddingResponse, error) {
	//payload := []byte(`{
	//	"model": "qwen2.5-vl-embedding",
	//	"input": {
	//		"contents": [
	//			{"text": "多模态向量模型"},
	//			{"image": "https://img.alicdn.com/imgextra/i3/O1CN01rdstgY1uiZWt8gqSL_!!6000000006071-0-tps-1970-356.jpg"},
	//			{"video": "https://help-static-aliyun-doc.aliyuncs.com/file-manage-files/zh-CN/20250107/lbcemt/new+video.mp4"},
	//			{"multi_images": [
	//				"https://img.alicdn.com/imgextra/i2/O1CN019eO00F1HDdlU4Syj5_!!6000000000724-2-tps-2476-1158.png",
	//				"https://img.alicdn.com/imgextra/i2/O1CN01dSYhpw1nSoamp31CD_!!6000000005089-2-tps-1765-1639.png"
	//			]}
	//		]
	//	}
	//}`)
	//多模态向量模型，这里注释是为了展现输入样式的，文本用了另外一个模型处理，所以这里只处理多模态
	payload := []byte(fmt.Sprintf(`{
		"model": "qwen2.5-vl-embedding",
		"input": {
			"contents": [
				{ %s: %s }
			]
		},
		"parameters": {
		"dimension":"2048"
}
	}`, ty, url))

	req, err := http.NewRequest(
		"POST", cfg.APIKey, bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}

	// 设置请求头
	req.Header.Set("Authorization", "Bearer "+"sk-737281fd1c884de7a57601f41c310e81")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取返回内容
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//return body, nil
	return nil, nil
}
