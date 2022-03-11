import { Button, Classes } from '@blueprintjs/core'
import { Popover2 } from '@blueprintjs/popover2'
import Image from 'next/image'
import { useContext, useEffect, useState } from 'react'
import { SubmitHandler, useForm } from 'react-hook-form'
import { ClientSocketMessage } from '../server/types'
import styles from '../styles/util.module.scss'
import { useWebsocket } from '../util/context'
import BottomMenu from './BottomMenu'
import { ReplyContext, SendData } from './ChatWindow'
import UserSettingsDialog from './UserSettingsDialog'

export const ChatBottomBar: React.VFC = (): JSX.Element => {
  const { connection, user } = useWebsocket()
  const [settingsOpen, setSettingsOpen] = useState(false)
  const [replyingTo, setReplyingTo] = useContext(ReplyContext)

  const { handleSubmit, register, reset, setFocus } = useForm<SendData>()

  const onSubmit: SubmitHandler<SendData> = (values) => {
    const message = values.message

    const socketMessage: ClientSocketMessage = {
      type: 'message',
      data: {
        content: message,
        parentId: replyingTo?.id,
      },
    }

    connection?.send(JSON.stringify(socketMessage))
    reset()
    setReplyingTo(undefined)
  }

  const loading = !user || !connection ? Classes.SKELETON : ''

  useEffect(() => {
    if (replyingTo) {
      setFocus('message')
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [replyingTo])

  useEffect(() => {
    const handler = (event: KeyboardEvent): void => {
      if (event.key === 'Escape') {
        event.preventDefault()
        setFocus('message')
      }
    }

    window.addEventListener('keydown', handler)

    return () => {
      window.removeEventListener('keydown', handler)
    }
  })

  return (
    <>
      {replyingTo && (
        <div className="flex flex-row items-center justify-between bg-darkgray-4">
          <div className="flex-1"></div>
          <span className="mr-2">Replying to</span>
          <Image
            src={replyingTo.author.avatarUrl ?? '/img/avatar.png'}
            width={16}
            height={16}
            className="rounded-full"
            alt="user avatar"
          />
          <b className="ml-1 mr-2">{replyingTo.author.username}</b>
          <Button minimal icon="cross" onClick={() => setReplyingTo(undefined)} />
          <div className="flex-1"></div>
        </div>
      )}
      <div className={'flex flex-row items-center justify-between h-[66px]'}>
        <div className="flex items-center justify-center h-full pl-2 bg-darkgray-3">
          <Popover2 content={<BottomMenu />} placement="top-start" minimal>
            <Button large minimal icon="menu" />
          </Popover2>
        </div>
        <div className="flex items-center h-full pl-3 bg-darkgray-3">
          <div
            className={`items-center rounded-full h-min ${styles['avatar-container']} ${loading} max-w-[250px]`}
          >
            <Image
              src={user?.avatarUrl ?? '/img/avatar.png'}
              width={40}
              height={40}
              className="rounded-full"
              alt="user avatar"
            />
          </div>
          <div className="ml-3">
            <div className={`text-lg font-semibold truncate ${loading}`}>
              {user?.username ?? 'a'.repeat(10)}
            </div>
            <div className={`text-xs truncate text-gray-3 ${loading}`}>
              {user?.id ?? 'a'.repeat(32)}
            </div>
          </div>
          <div className="ml-3">
            <Button
              minimal
              large
              icon="cog"
              className={loading}
              onClick={() => setSettingsOpen(true)}
            />
            <UserSettingsDialog isOpen={settingsOpen} setOpen={setSettingsOpen} />
          </div>
        </div>
        <div className="flex flex-row items-center justify-center flex-grow h-full bg-darkgray-3">
          <form onSubmit={handleSubmit(onSubmit)} className="flex w-full h-full px-4 py-1">
            <input
              className={`w-full my-2 pl-4 pr-12 rounded-lg bg-darkgray-5 ${loading}`}
              placeholder={!loading ? 'Enter message...' : undefined}
              {...register('message', { required: true })}
            />
            <Button
              minimal
              large
              type="submit"
              icon="send-message"
              className={`w-12 my-2 -ml-12 rounded-l-none rounded-r-lg ${loading}`}
            />
          </form>
        </div>
      </div>
    </>
  )
}
