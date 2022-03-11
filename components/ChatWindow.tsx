import React, { createContext, useEffect, useRef, useState } from 'react'
import { ServerMessage, ServerSocketMessage } from '../server/types'
import { useMessages, useWebsocket } from '../util/context'
import { ChatBottomBar } from './ChatBottomBar'
import ChatMessage from './ChatMessage'
import Loading from './Loading'
import SimpleChatMessage from './SimpleChatMessage'

export interface SendData {
  message: string
}

export type ReplyContextValue = [
  ServerMessage['data'] | undefined,
  React.Dispatch<React.SetStateAction<ServerMessage['data'] | undefined>>
]
export const ReplyContext = createContext<ReplyContextValue>([undefined, () => undefined])

const ChatWindow: React.VFC = () => {
  const { connection, user } = useWebsocket()
  const [messages, setMessages] = useMessages()

  const [replyingTo, setReplyingTo] = useState<ServerMessage['data'] | undefined>()

  const messageContainer = useRef<HTMLDivElement>(null)

  useEffect(() => {
    if (!connection) return

    const handler = (event: MessageEvent): void => {
      const message = JSON.parse(event.data) as ServerSocketMessage
      if (message.type === 'message') {
        setMessages((messages) => [...messages, message.data as ServerMessage['data']])
        if (messageContainer.current) {
          messageContainer.current.scrollTop = messageContainer.current.scrollHeight
        }
      }
    }

    connection.addEventListener('message', handler)

    const fetchMessages = async (): Promise<void> => {
      const res = await fetch(
        `https://${process.env.API_HOSTNAME}:${process.env.API_PORT}/c/public`
      )
      const json = await res.json()
      // eslint-disable-next-line @typescript-eslint/no-unsafe-member-access
      const messages = json.data.messages as Array<ServerMessage['data']>
      messages.sort((a, b) => a.createdAt - b.createdAt)
      setMessages(messages)
      if (messageContainer.current) {
        messageContainer.current.scrollTop = messageContainer.current.scrollHeight
      }
    }
    void fetchMessages()

    return () => {
      connection.removeEventListener('message', handler)
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [connection])

  if (!connection || !user) {
    return <Loading />
  }

  return (
    <ReplyContext.Provider value={[replyingTo, setReplyingTo]}>
      <div className="flex flex-col w-full h-full">
        <div
          id="messages"
          className="justify-end flex-grow pb-4 overflow-y-scroll"
          ref={messageContainer}
        >
          {messages.map((message, i) => {
            if (
              i === 0 ||
              message.createdAt - messages[i - 1].createdAt >= 1000 * 60 * 10 ||
              message.author.id !== messages[i - 1].author.id ||
              message.parentId
            ) {
              return (
                <React.Fragment key={i}>
                  <div className={i === 0 ? 'h-4' : 'h-2'}></div>
                  <ChatMessage message={message} key={message.id} />
                </React.Fragment>
              )
            } else {
              return <SimpleChatMessage message={message} key={message.id} />
            }
          })}
        </div>
        <ChatBottomBar />
      </div>
    </ReplyContext.Provider>
  )
}

export default ChatWindow
