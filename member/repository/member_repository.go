package repository

import (
	"context"
	"fmt"
	requestDto "github.com/YooGenie/daily-work-log-service/dto/request"
	"github.com/YooGenie/daily-work-log-service/member/entity"
	"sync"
)

var (
	memberRepositoryOnce     sync.Once
	memberRepositoryInstance *memberRepository
)

func MemberRepository() *memberRepository {
	memberRepositoryOnce.Do(func() {
		memberRepositoryInstance = &memberRepository{}
	})
	return memberRepositoryInstance
}

type memberRepository struct {
}

func (memberRepository) Create(ctx context.Context, creation requestDto.MemberCreate) error {

	member := entity.Member{
		Email:    creation.Email,
		Password: creation.Password,
		Name:     creation.Name,
		Mobile:   creation.Mobile,
	}

	fmt.Println(member)

	//echo에서 context 만들기

	return nil
}
