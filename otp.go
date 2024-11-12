package main

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type OTP struct {	
	Key     string
	Created time.Time
}

type RetentionMap map[string]OTP

func NewRetentionMap(ctx context.Context, retentionPeriod time.Duration) RetentionMap {
	rm := make(RetentionMap)
	go rm.Retention(ctx, retentionPeriod)
	return rm
}

func (m RetentionMap) NewOTP() OTP {
	o := OTP{
		Key:     uuid.NewString(),
		Created: time.Now(),
	}
	m[o.Key] = o
	return o
}

func (rm RetentionMap) VerifyOTP(otp string) bool {
	_, ok := rm[otp]
	if !ok {
		return false
	}
	// delete as it is "one time password" so delete it as soon as used
	delete(rm, otp)
	return true
}

func (rm RetentionMap) Retention(ctx context.Context, retentionPeriod time.Duration) {
	ticker := time.NewTicker(400 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			for _, otp := range rm {
				if otp.Created.Add(retentionPeriod).Before(time.Now()) {
					delete(rm, otp.Key)
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
