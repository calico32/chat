import { Button, ButtonGroup } from '@blueprintjs/core'
import React, { useContext } from 'react'
import { ServerMessage } from '../server/types'
import { ReplyContext } from './ChatWindow'

interface ChatMessageActionsProps {
  message: ServerMessage['data']
  className?: string
}

const ChatMessageActions: React.VFC<ChatMessageActionsProps> = ({ message, className }) => {
  const [, setReplyingTo] = useContext(ReplyContext)

  return (
    <ButtonGroup className={`absolute -top-3 right-2 float-right z-20 ${className}`}>
      <Button icon="undo" onClick={() => setReplyingTo(message)} className="z-20" />
    </ButtonGroup>
  )
}

export default ChatMessageActions
