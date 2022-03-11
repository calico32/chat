export interface User {
  id: string
  username: string
  avatarUrl?: string
  secretHash?: string
}

export interface SocketMessage {
  type: string
  data: unknown
}

export type ClientSocketMessage = ClientLogin | ClientMessage | ClientPrivateMessage
export type ServerSocketMessage =
  | ServerLogin
  | ServerError
  | ServerMessage
  | ServerPrivateMessage
  | ServerClear
  | ServerReload

export interface ClientLogin extends SocketMessage {
  type: 'login'
  data: {
    user: {
      username: string
      avatarUrl?: string
      secretHash?: string
    }
  }
}

export interface ClientMessage extends SocketMessage {
  type: 'message'
  data: {
    content: string
    parentId?: string
  }
}

export interface ClientPrivateMessage extends SocketMessage {
  type: 'privateMessage'
  data: {
    content: string
    toId: string
    parentId?: string
  }
}

export interface ServerLogin extends SocketMessage {
  type: 'login'
  data: {
    user: User
  }
}

export interface ServerError extends SocketMessage {
  type: 'error'
  data: {
    error: string
  }
}

export interface ServerMessage extends SocketMessage {
  type: 'message'
  data: {
    id: string
    content: string
    parentId?: string
    author: User
    /** Milliseconds since Unix epoch */
    createdAt: number
  }
}

export interface ServerPrivateMessage extends SocketMessage {
  type: 'message'
  data: {
    id: string
    content: string
    from: User
    parentId?: string
  }
}

export interface ServerClear extends SocketMessage {
  type: 'clear'
  data: null
}

export interface ServerReload extends SocketMessage {
  type: 'reload'
  data: null
}

// http

export interface HttpMessage extends SocketMessage {}

export type ServerHttpMessage =
  | ServerMessageList
  | ServerPrivateMessageList
  | ServerPrivateMessageRecipientList

export interface ServerMessageList extends HttpMessage {
  type: 'messages'
  data: {
    beforeId?: string
    afterId?: string
    count: number
    messages: ServerMessage[]
  }
}

export interface ServerPrivateMessageList extends HttpMessage {
  type: 'private_messages'
  data: {
    beforeId?: string
    afterId?: string
    count: number
    messages: ServerPrivateMessage[]
  }
}

export interface ServerPrivateMessageRecipientList extends HttpMessage {
  type: 'private_message_recipients'
  data: {
    recipients: User[]
  }
}
