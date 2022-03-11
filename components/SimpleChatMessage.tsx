import React, { useContext } from 'react'
import { ServerMessage } from '../server/types'
import styles from '../styles/chatmessage.module.scss'
import { getLocale } from '../util/util'
import ChatMessageActions from './ChatMessageActions'
import { ReplyContext } from './ChatWindow'

interface SimpleChatMessageProps {
  message: ServerMessage['data']
}

const SimpleChatMessage: React.VFC<SimpleChatMessageProps> = ({ message }) => {
  const [replyingTo] = useContext(ReplyContext)

  if (message.parentId) {
    console.warn(
      'SimpleChatMessage: a message has a parentId, this is not supported in SimpleChatMessage'
    )
  }

  const time = new Date(message.createdAt)
  const formattedTime = time.toLocaleTimeString(getLocale(), {
    hour: 'numeric',
    minute: '2-digit',
  })

  const beingRepliedTo = replyingTo?.id === message.id

  return (
    <div
      className={`flex py-px relative items-center w-full flex-nowrap max-w-full hover:backdrop-brightness-125 ${
        styles.simple
      } ${beingRepliedTo ? 'bg-cobalt-1 !backdrop-brightness-0 border-l-cobalt-3 border-l-2' : ''}`}
    >
      <ChatMessageActions message={message} className={styles.actions} />
      <span
        className={`text-[0.6rem] [line-height:0.5rem] text-center text-gray-3 ${
          styles.timestamp
        } ${beingRepliedTo ? 'min-w-[62px] w-[62px]' : 'min-w-[64px] w-[64px]'}`}
      >
        {formattedTime}
      </span>
      <div className="mr-3 break-words overflow-x-clip">
        <span className="break-words">{message.content}</span>
      </div>
    </div>
  )
}

export default SimpleChatMessage
