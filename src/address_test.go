package src_test

import (
	"testing"
	"dy201.com/addressTrans/src"
)

func TestGetAddressJSON(t *testing.T)  {
	address,err := src.GetAddressJSON("北京市海淀区上地十街10号")
	if err != nil {
		t.Error("获取地址失败，",err)
	}

	if address.Result.Location.Lng != 116.3084202915042 {
		t.Error("经度计算错误")
	}

	if address.Result.Location.Lat != 40.05703033345938 {
		t.Error("维度计算错误")
	}


}