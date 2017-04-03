package consistent_hashing



type ConsistentHashingConfig struct {
	hashRange uint32
}


func NewConsistentHashingConfig(hashRange uint32) *ConsistentHashingConfig {
	config := &ConsistentHashingConfig{hashRange}
	return config
}

func (config *ConsistentHashingConfig) getHashRange() uint32{
	return config.hashRange
}