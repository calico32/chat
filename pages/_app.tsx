/* eslint-disable @typescript-eslint/no-invalid-void-type */
import { FocusStyleManager, Toaster } from '@blueprintjs/core'
import type { AppProps } from 'next/app'
import { useEffect, useState } from 'react'
import { ServerMessage, ServerSocketMessage, User } from '../server/types'
import '../styles/globals.scss'
import {
  localStorageManager,
  MessagesContext,
  ToasterContext,
  WebsocketContext,
  WebsocketState,
} from '../util/context'
import { sendMessage } from '../util/user'

const toaster = typeof window !== 'undefined' ? Toaster.create({ position: 'top-right' }) : null

function App({ Component, pageProps }: AppProps): JSX.Element {
  const [messages, setMessages] = useState<Array<ServerMessage['data']>>([])
  const [websocketState, setWebsocketState] = useState<WebsocketState>({
    connection: null,
    user: null,
    updateUser: (update: User | ((user: User | null) => User)): void => {
      const newUser = typeof update === 'function' ? update(websocketState.user) : update
      if (
        newUser.username !== websocketState.user?.username ||
        newUser.secretHash !== websocketState.user?.secretHash
      ) {
        window.location.reload()
      }
    },
  })

  useEffect(() => {
    FocusStyleManager.onlyShowFocusOnTabs()

    const ws = new WebSocket(`wss://${process.env.API_HOSTNAME}:${process.env.API_PORT}/ws`)

    const onOpen = (): void => {
      sendMessage(ws, 'login', {
        user: {
          username: localStorageManager.username ?? randomUsername(),
          avatarUrl: localStorageManager.avatarUrl ?? undefined,
          secretHash: localStorageManager.secretHash ?? undefined,
        },
      })
    }

    const onMessage = (event: MessageEvent<any>): void => {
      const message = JSON.parse(event.data) as ServerSocketMessage
      switch (message.type) {
        case 'login':
          setWebsocketState({
            connection: ws,
            user: message.data.user,
            updateUser: websocketState.updateUser,
          })
          break
        case 'error':
          toaster?.show({
            icon: 'error',
            intent: 'danger',
            message: `Server error: ${message.data.error}`,
          })
          break
        case 'clear':
          setMessages([])
          break
        case 'reload':
          window.location.reload()
          break
      }
    }

    ws.addEventListener('open', onOpen)
    ws.addEventListener('message', onMessage)

    return () => {
      ws.removeEventListener('open', onOpen)
      ws.removeEventListener('message', onMessage)
    }

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  return (
    <ToasterContext.Provider value={toaster!}>
      <MessagesContext.Provider value={[messages, setMessages]}>
        <WebsocketContext.Provider value={websocketState}>
          <Component {...pageProps} />
        </WebsocketContext.Provider>
      </MessagesContext.Provider>
    </ToasterContext.Provider>
  )
}

export default App

const randomUsername = (): string => `user-${crypto.getRandomValues(new Uint8Array(4)).join('')}`
