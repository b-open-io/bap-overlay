package bap

import (
	"context"
	"encoding/json"

	"github.com/b-open-io/overlay/storage"
	"github.com/redis/go-redis/v9"
)

type BAPStorage struct {
	*storage.RedisStorage
}

func (s *BAPStorage) LoadIdentityById(ctx context.Context, id string) (*Identity, error) {
	if b, err := s.DB.HGet(ctx, IdentityKey, id).Bytes(); err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {
		var identity Identity
		if err := json.Unmarshal(b, &identity); err != nil {
			return nil, err
		}
		return &identity, nil
	}
}

func (s *BAPStorage) LoadIdentityByAddress(ctx context.Context, address string) (*Identity, error) {
	if idkey, err := s.DB.HGet(ctx, AddressIdentityKey, address).Result(); err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else if b, err := s.DB.HGet(ctx, IdentityKey, idkey).Bytes(); err != nil {
		return nil, err
	} else {
		var identity Identity
		if err := json.Unmarshal(b, &identity); err != nil {
			return nil, err
		}
		return &identity, nil
	}
}

func (s *BAPStorage) SaveIdentity(ctx context.Context, id *Identity) error {
	if b, err := json.Marshal(id); err != nil {
		return err
	} else {
		_, err := s.DB.Pipelined(ctx, func(p redis.Pipeliner) error {
			if err := p.HSet(ctx, IdentityKey, id.IDKey, b).Err(); err != nil {
				return err
			} else if err := p.HSet(ctx, AddressIdentityKey, id.CurrentAddress, id.IDKey).Err(); err != nil {
				return err
			}

			return nil
		})
		return err
	}
}

func (s *BAPStorage) LoadAttestation(ctx context.Context, urnHash string) (*Attestation, error) {
	if b, err := s.DB.HGet(ctx, AttestationKey, urnHash).Bytes(); err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {
		var att Attestation
		if err := json.Unmarshal(b, &att); err != nil {
			return nil, err
		}
		return &att, nil
	}
}

func (s *BAPStorage) SaveAttestation(ctx context.Context, att *Attestation) error {
	if b, err := json.Marshal(att); err != nil {
		return err
	} else {
		return s.DB.HSet(ctx, AttestationKey, att.Id, b).Err()
	}
}

func (s *BAPStorage) LoadProfile(ctx context.Context, bapId string) (json.RawMessage, error) {
	if msg, err := s.DB.HGet(ctx, ProfileKey, bapId).Bytes(); err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {
		return msg, nil
	}
}

func (s *BAPStorage) SaveProfile(ctx context.Context, bapId string, profile json.RawMessage) error {
	return s.DB.HSet(ctx, ProfileKey, bapId, profile).Err()
}
