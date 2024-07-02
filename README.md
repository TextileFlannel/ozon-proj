```docker-compose up --build```

## Реализованный функционал:
1) Создание поста (есть возможность отключить комментарии)
```
mutation {
  createPost(input: { title: "qwe", content: "123", commentsAllowed: true }) {
    id
    title
    content
    commentsAllowed
  }
}
```
2) Создание комментария (!!!Для корректного обращения к посту лучше вывести их все, тк postID генерируется случайно)
```
mutation {
  createComment(input: { postID: "1", content: "qweqwe" }) {
    id
    postID
    content
    createdAt
  }
}
```
3) Вывод всех постов
```
query {
  posts {
    id
    title
    content
    commentsAllowed
  }
}
```
4) Вывод поста по Id с комментариями
```
query {
  post(id: "1") {
    id
    title
    content
    comments {
      id
      content
      createdAt
    }
  }
}
```
5) Вывод всех постов и комментариев
```
query {
  posts {
    id
    title
    content
    commentsAllowed
    comments {
      id
      content
    }
  }
}
```
