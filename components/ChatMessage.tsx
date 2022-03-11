import Image from 'next/image'
import React from 'react'
import { ServerMessage } from '../server/types'
import styles from '../styles/chatmessage.module.scss'
import utilStyles from '../styles/util.module.scss'
import { useMessages } from '../util/context'
import { getLocale, truncate } from '../util/util'
import ChatMessageActions from './ChatMessageActions'
import UserCard from './UserCard'

interface ChatMessageProps {
  message: ServerMessage['data']
}

const ChatMessage: React.VFC<ChatMessageProps> = ({ message }) => {
  const time = new Date(message.createdAt)
  const formattedTime = time.toLocaleString(getLocale(), {})
  const [messages] = useMessages()

  const parent = message.parentId ? messages.find((m) => m.id === message.parentId) : undefined

  return (
    <div className={`relative flex flex-col hover:backdrop-brightness-125 ${styles.message}`}>
      <ChatMessageActions message={message} className={styles.actions} />
      {message.parentId &&
        (!parent ? (
          <div className="flex">
            <div className="w-[32px]" />
            <span className="italic">Message could not be loaded.</span>
          </div>
        ) : (
          <div className={`relative flex items-center text-xs font-light ${styles.replyContext}`}>
            <div className="w-[64px]" />
            <div className={styles.replyIndicator} />
            <Image
              src={parent.author.avatarUrl ?? '/img/avatar.png'}
              width={16}
              height={16}
              className="rounded-full"
              alt="user avatar"
            />
            <span className="ml-1">@{parent.author.username}</span>
            <span className={`ml-1 text-gray-4 ${styles.replyContent}`}>
              {truncate(parent.content, 100)}
            </span>
          </div>
        ))}
      <div className="flex px-4 py-px">
        <div className="flex items-center justify-center">
          <UserCard
            user={message.author}
            target={({ ref, ...targetProps }) => (
              <div ref={ref} {...targetProps} className={utilStyles.avatarContainer}>
                <Image
                  src={message.author.avatarUrl ?? '/img/avatar.png'}
                  width={36}
                  height={36}
                  className="transition-all rounded-full cursor-pointer max-h-12 hover:brightness-50 active:translate-y-px"
                  alt="user avatar"
                />
              </div>
            )}
          />
        </div>
        <div className="flex flex-col justify-between ml-3">
          <div className="flex items-baseline">
            <UserCard
              user={message.author}
              target={({ ref, ...targetProps }) => (
                <span
                  ref={ref}
                  {...targetProps}
                  className="font-bold cursor-pointer gerald hover:underline"
                >
                  {message.author.username}
                </span>
              )}
            />
            <span className="ml-2 text-gray-1">{formattedTime}</span>
          </div>
          <span>{message.content}</span>
        </div>
      </div>
    </div>
  )
}

export default ChatMessage
