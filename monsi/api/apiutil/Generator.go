package apiutil

import (
	"monsi/api/model"
	"monsi/util"
	"monsi/vcmanager"
)

// Generates a Reviewed VC object which is invalid
func GenReviewedVC(vc util.VC) model.ReviewedVCDTO {
	var rvc model.ReviewedVCDTO

	rvc.ID = vc.ID
	rvc.Type = vc.Type
	rvc.Subject = vc.Subject
	rvc.MonsiValid = vcmanager.CheckValidityOfVC(vc)

	return rvc
}
