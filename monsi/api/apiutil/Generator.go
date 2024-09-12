package apiutil

import (
	"monsi/api/model"
	"monsi/util"
	"monsi/vcmanager"
)

// Generates a Reviewed VC object which is invalid
func GenReviewedVC(vc util.VC) model.ReviewedVCDTO {
	var rvc model.ReviewedVCDTO

	rvc.Context = vc.Context
	rvc.ID = vc.ID
	rvc.Type = vc.Type
	rvc.Issuer = vc.Issuer
	rvc.ValidFrom = vc.ValidFrom
	rvc.ValidUntil = vc.ValidUntil
	rvc.Subject = vc.Subject
	rvc.Proof = vc.Proof
	rvc.MonsiValid = vcmanager.CheckValidityOfVC(vc)

	return rvc
}
