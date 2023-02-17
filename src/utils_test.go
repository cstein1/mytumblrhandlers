package mytumblrhandlers

import (
	"fmt"
	"testing"
)

func TestExtractTextHTML(t *testing.T) {
	str, err := ExtractTextHTML(testHtml)
	if err != nil {
		t.Errorf("failed to extract html: %v", err.Error())
	}
	fmt.Println(str)
	if str != correctHtml {
		t.Fatal("affected string does not match content")
	}
}

var testHtml string = `<p><figure class="tmblr-full"><img src="https://64.media.tumblr.com/73d2d368bdb36945bb9efd2ba59699d0/ffb59d4b7ef5aac2-84/s640x960/787bfc394499f3caa2f4942990ba492b35c14ca5.jpg" alt="image" class=""/></figure><p>nope! here are her baby pictures (from 2017)</p><figure class="tmblr-full"><img src="https://64.media.tumblr.com/d908ff6a36c53eddde3a40cb62176238/ffb59d4b7ef5aac2-30/s640x960/ea970d8f70c1b58f140de4244cd2464bf528c304.jpg" alt="image" class=""/></figure><figure class="tmblr-full"><img src="https://64.media.tumblr.com/a9e97d2e1f2719b976d660e6e19b1042/ffb59d4b7ef5aac2-35/s640x960/766d99a58db78ff30618f7860b642fca9f0e2222.jpg" alt="image" class=""/></figure><figure class="tmblr-full"><img src="https://64.media.tumblr.com/c9c3336c9cd390636dcf2dd5c389fb82/ffb59d4b7ef5aac2-13/s640x960/111677c5f3483e3bb59b8fe7c9797725ef9aa9a8.jpg" alt="image" class=""/></figure><figure class="tmblr-full"><img src="https://64.media.tumblr.com/efaa9313dea9dc50b9354ca12b64fc2e/ffb59d4b7ef5aac2-cc/s640x960/eda487d778d6147b7426f84884eb6c5e4e333911.jpg" alt="image" class=""/></figure><p>she went through an almost normal cat phase around 3 months of age but reverted back to being yucky</p></p>`
var correctHtml string = " https://64.media.tumblr.com/73d2d368bdb36945bb9efd2ba59699d0/ffb59d4b7ef5aac2-84/s640x960/787bfc394499f3caa2f4942990ba492b35c14ca5.jpg nope! here are her baby pictures (from 2017)\n https://64.media.tumblr.com/d908ff6a36c53eddde3a40cb62176238/ffb59d4b7ef5aac2-30/s640x960/ea970d8f70c1b58f140de4244cd2464bf528c304.jpg  https://64.media.tumblr.com/a9e97d2e1f2719b976d660e6e19b1042/ffb59d4b7ef5aac2-35/s640x960/766d99a58db78ff30618f7860b642fca9f0e2222.jpg  https://64.media.tumblr.com/c9c3336c9cd390636dcf2dd5c389fb82/ffb59d4b7ef5aac2-13/s640x960/111677c5f3483e3bb59b8fe7c9797725ef9aa9a8.jpg  https://64.media.tumblr.com/efaa9313dea9dc50b9354ca12b64fc2e/ffb59d4b7ef5aac2-cc/s640x960/eda487d778d6147b7426f84884eb6c5e4e333911.jpg she went through an almost normal cat phase around 3 months of age but reverted back to being yucky\n"

func TestConvertTumblrTimeToEpoch(t *testing.T) {
	sampleTime := "2023-02-18 17:12:38 GMT"
	timeStr, err := ConvertTumblrTimeToEpoch(sampleTime)
	if err != nil {
		t.Fatalf("failed to convert with err: %v", err.Error())
	}
	fmt.Printf("%s", timeStr)
	if timeStr != "1676740358" {
		t.Fatal("did not convert to proper epoch")
	}
}
