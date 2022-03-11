import { IToaster } from '@blueprintjs/core'
import { createContext, useContext } from 'react'
import { ServerMessage, User } from '../server/types'

export const ToasterContext = createContext<IToaster>(undefined!)

export const useToaster = (): IToaster => {
  const toaster = useContext(ToasterContext)
  return toaster
}

export interface WebsocketState {
  connection: WebSocket | null
  user: User | null
  updateUser: ((user: User) => void) & ((updater: (user: User | null) => User) => void)
}

export const WebsocketContext = createContext<WebsocketState>(undefined!)

export const useWebsocket = (): WebsocketState => {
  const socket = useContext(WebsocketContext)
  return socket
}

export const localStorageManager = {
  namespace: 'chat',
  getItem(key: string) {
    return localStorage.getItem(`${this.namespace}:${key}`)
  },
  setItem(key: string, value: string | null) {
    if (value) {
      localStorage.setItem(`${this.namespace}:${key}`, value)
    } else {
      localStorage.removeItem(`${this.namespace}:${key}`)
    }
  },

  get username(): string | null {
    return this.getItem('username')
  },
  set username(username: string | null) {
    this.setItem('username', username)
  },

  get avatarUrl(): string | null {
    return this.getItem('avatarUrl')
  },
  set avatarUrl(avatarUrl: string | null) {
    this.setItem('avatarUrl', avatarUrl)
  },

  get secretHash(): string | null {
    return this.getItem('secretHash')
  },
  set secretHash(secretHash: string | null) {
    this.setItem('secretHash', secretHash)
  },

  clear() {
    this.username = null
    this.avatarUrl = null
    this.secretHash = null
  },
}

export type MessagesState = [
  messages: Array<ServerMessage['data']>,
  setMessages: React.Dispatch<React.SetStateAction<Array<ServerMessage['data']>>>
]

export const MessagesContext = createContext<MessagesState>(undefined!)

export const useMessages = (): MessagesState => {
  const messages = useContext(MessagesContext)
  return messages
}
