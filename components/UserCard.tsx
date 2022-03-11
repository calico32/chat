import { IPopoverProps } from '@blueprintjs/core'
import { Popover2, Popover2TargetProps } from '@blueprintjs/popover2'
import Image from 'next/image'
import React from 'react'
import { User } from '../server/types'
import { useWebsocket } from '../util/context'

interface UserCardProps {
  user: User
  target: (props: Popover2TargetProps & IPopoverProps) => JSX.Element
}

const UserCard: React.FC<UserCardProps> = ({ target, user }) => {
  const { user: currentUser } = useWebsocket()

  return (
    <Popover2
      placement="auto-start"
      minimal
      content={
        <div className="flex flex-col p-4 text-center">
          <div className="flex items-center justify-center">
            <Image
              src={user.avatarUrl ?? '/img/avatar.png'}
              width={64}
              height={64}
              className="rounded-full"
              alt="user avatar"
            />
          </div>
          <span className="mt-2 text-lg font-bold gerald">{user.username}</span>
          <span className="mt-px text-sm text-gray-3">{user.id}</span>
          {user.id === currentUser?.id && <span className="mt-1 font-bold text-gray-4">You</span>}
        </div>
      }
      renderTarget={target}
    ></Popover2>
  )
}

export default UserCard
