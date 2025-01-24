package builder

import (
	relativesexternalrpc "github.com/PhuPhuoc/curanest-patient-service/module/relatives/infars/externalrpc"
	relativesrepository "github.com/PhuPhuoc/curanest-patient-service/module/relatives/infars/repository"
	relativescommands "github.com/PhuPhuoc/curanest-patient-service/module/relatives/usecase/commands"
	relativesqueries "github.com/PhuPhuoc/curanest-patient-service/module/relatives/usecase/queries"
	"github.com/jmoiron/sqlx"
)

type builderOfRelatives struct {
	db                    *sqlx.DB
	urlPathAccountService string
}

func NewRelativesBuilder(db *sqlx.DB) builderOfRelatives {
	return builderOfRelatives{db: db}
}

func (s builderOfRelatives) AddUrlPathAccountService(url string) builderOfRelatives {
	s.urlPathAccountService = url
	return s
}

func (s builderOfRelatives) BuildExternalAccountServiceInCmd() relativescommands.ExternalAccountService {
	return relativesexternalrpc.NewAccountRPC(s.urlPathAccountService)
}

func (s builderOfRelatives) BuildRelativesCmdRepo() relativescommands.RelativeCommandRepo {
	return relativesrepository.NewRelativesRepo(s.db)
}

func (s builderOfRelatives) BuildRelativesQueryRepo() relativesqueries.RelativesQueryRepo {
	return relativesrepository.NewRelativesRepo(s.db)
}

func (s builderOfRelatives) BuildExternalAccountServiceInQuery() relativesqueries.ExternalAccountService {
	return relativesexternalrpc.NewAccountRPC(s.urlPathAccountService)
}
