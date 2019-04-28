package server

// Config is configure for server node
type Config struct {
	port                 uint16
	coordinatorAddresses []string

	crReplicateFactor         uint8
	crReplicateAfterAggregate bool

	badgerDir string
}
