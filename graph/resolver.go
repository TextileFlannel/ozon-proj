package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
    "time"
    "context"

    "github.com/TextileFlannel/ozon-proj/graph/model"
    "gorm.io/gorm"
    "github.com/google/uuid"
)

type Resolver struct{
	DB *gorm.DB
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
    post := &model.Post{
        ID:              uuid.New().String(),
		Title:           input.Title,
		Content:         input.Content,
		CommentsAllowed: input.CommentsAllowed,
		CreatedAt:       time.Now().Format(time.RFC3339),
	}
    if err := r.DB.Create(post).Error; err != nil {
        return nil, err
    }
    return post, nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error) {
    comment := &model.Comment{
        ID:        uuid.New().String(),
		PostID:    input.PostID,
		ParentID:  input.ParentID,
		Content:   input.Content,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
    if err := r.DB.Create(comment).Error; err != nil {
        return nil, err
    }
    return comment, nil
}

func (r *Resolver) Mutation() MutationResolver {
    return &mutationResolver{r}
}

type queryResolver struct{*Resolver}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
    var posts []*model.Post
    if err := r.DB.Preload("Comments").Find(&posts).Error; err != nil {
        return nil, err
    }
    return posts, nil
}

func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
    var post model.Post
    if err := r.DB.Preload("Comments").First(&post, "id=?", id).Error; err != nil {
        return nil, err
    }
    return &post, nil
}

func (r *Resolver) Query() QueryResolver {
    return &queryResolver{r}
}