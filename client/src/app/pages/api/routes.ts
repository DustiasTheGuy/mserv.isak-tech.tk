export const routes = [{
    path: '/api/posts',
    method: 'GET',
    data: null
  },
  {
    path: '/api/post/:id',
    method: 'GET',
    data: null
  },
  {
    path: '/api/paginate/:page/:limit',
    method: 'GET',
    data: null
  },
  {
    path: '/api/delete',
    method: 'DELETE',
    data: {
      "id": "uint64"
    }
  },
  {
    path: '/api/new',
    method: 'POST',
    data: {
      "title": "string",
      "body": "string",
      "tags": "[]string"
    }
  },
  {
    path: '/api/update',
    method: 'PUT',
    data: {
      "id": "uint64",
      "title": "string",
      "body": "string",
      "tags": "[]string"
    }
  }
]
