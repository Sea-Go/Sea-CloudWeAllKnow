package infra

import (
	"context"
	"sea/config"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func Neo4jInit() error {
	cfg := config.Cfg
	ctx := context.Background()
	client, err := neo4j.NewDriverWithContext(

		cfg.Neo4j.Address,
		neo4j.BasicAuth(
			cfg.Neo4j.Username,
			cfg.Neo4j.Password,
			"",
		),
	)
	if err != nil {
		return err
	}
	defer client.Close(ctx)
	err = client.VerifyConnectivity(ctx)
	if err != nil {
		return err
	}
	return nil
}
