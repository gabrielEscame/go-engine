package engine

import (
	"github.com/gabrielEscame/go-engine/physics"
)

type CollisionInfo struct {
	EntityCollided *CollidableEntity
}

func CollisionChecker(entities []CollidableEntity) {
	for i, a := range entities {
		for j := i + 1; j < len(entities); j++ {
			b := entities[j]

			if physics.AABB(a.GetShape(), b.GetShape()) {
				a.OnCollisionEnter(NewCollisionInfo(&b))
				b.OnCollisionEnter(NewCollisionInfo(&a))
			}
		}
	}
}

func NewCollisionInfo(entityCollided *CollidableEntity) *CollisionInfo {
	return &CollisionInfo{
		EntityCollided: entityCollided,
	}
}
