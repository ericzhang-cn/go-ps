package server

// ServerConfig is configure for server node
type ServerConfig struct {
	port                 uint16
	coordinatorAddresses []string

	crReplicateFactor         uint8
	crReplicateAfterAggregate bool

	boltDBName string
}
