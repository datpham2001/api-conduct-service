package model

import (
	"time"

	"github.com/datpham2001/realworld-conduct/model/enum"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID              *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	CreatedTime     *time.Time          `json:"createdTime,omitempty" bson:"created_time,omitempty"`
	LastUpdatedTime *time.Time          `json:"lastUpdatedTime,omitempty" bson:"last_updated_time,omitempty"`

	ArticleID            string `json:"articleId,omitempty" bson:"article_id,omitempty"`
	AuthorID             string `json:"authorId,omitempty" bson:"author_id,omitempty"`
	Title                string `json:"title,omitempty" bson:"title,omitempty"`
	Description          string `json:"description,omitempty" bson:"description,omitempty"`
	Slug                 string `json:"slug,omitempty" bson:"slug,omitempty"`
	Body                 string `json:"body,omitempty" bson:"body,omitempty"`
	EstimatedTimeReading int64  `json:"estimatedTimeReading,omitempty" bson:"estimated_time_reading,omitempty"` // by minute/hour
	IsFavorited          *bool  `json:"isFavorited,omitempty" bson:"is_favorited,omitempty"`
	FavoritesCount       int64  `json:"favoritesCount,omitempty" bson:"favorites_count,omitempty"`

	ComplexQuery []*bson.M `json:"-" bson:"$and,omitempty"`
}

type Hashtag struct {
	ID              *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	CreatedTime     *time.Time          `json:"createdTime,omitempty" bson:"created_time,omitempty"`
	LastUpdatedTime *time.Time          `json:"lastUpdatedTime,omitempty" bson:"last_updated_time,omitempty"`

	HashtagID   string `json:"hashtagId,omitempty" bson:"hashtag_id,omitempty"`
	HashtagName string `json:"hashtagName,omitempty" bson:"hashtag_name,omitempty"`

	ComplexQuery []*bson.M `json:"-" bson:"$and,omitempty"`
}

type ArticleHashtag struct {
	ID              *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	CreatedTime     *time.Time          `json:"createdTime,omitempty" bson:"created_time,omitempty"`
	LastUpdatedTime *time.Time          `json:"lastUpdatedTime,omitempty" bson:"last_updated_time,omitempty"`

	ArticleID string `json:"articleId,omitempty" bson:"article_id,omitempty"`
	HashtagID string `json:"hashtagId,omitempty" bson:"hashtag_id,omitempty"`
}

type Comment struct {
	ID              *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	CreatedTime     *time.Time          `json:"createdTime,omitempty" bson:"created_time,omitempty"`
	LastUpdatedTime *time.Time          `json:"lastUpdatedTime,omitempty" bson:"last_updated_time,omitempty"`

	ParentCommentID string   `json:"parentCommentId,omitempty" bson:"parent_comment_id,omitempty"`
	CommentID       string   `json:"commentId,omitempty" bson:"comment_id,omitempty"`
	AuthorID        string   `json:"authorId,omitempty" bson:"author_id,omitempty"`
	ArticleID       string   `json:"articleId,omitempty" bson:"article_id,omitempty"`
	Body            string   `json:"body,omitempty" bson:"body,omitempty"`
	ReactObjects    []*React `json:"reactObjects,omitempty" bson:"react_objects,omitempty"`
}

type React struct {
	ReactType  enum.ReactTypeValue `json:"reactType,omitempty" bson:"react_type,omitempty"`
	ReactCount int64               `json:"reactCount,omitempty" bson:"react_account,omitempty"`
	ReactBy    []*string           `json:"reactBy,omitempty" bson:"react_by,omitempty"`
}
