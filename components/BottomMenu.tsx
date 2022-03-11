import { Button, Callout, Colors, H6, Icon } from '@blueprintjs/core'
import React from 'react'

const BottomMenu: React.VFC = () => {
  const canSendPrivateMessages = false // user?.secretHash != null

  return (
    <div className="max-w-xs p-4">
      <H6 className="tracking-wider uppercase !text-gray-3">Public Channels</H6>
      <span className="flex items-center ml-2">
        <Icon icon="chat" color={Colors.GRAY3} className="mr-2" />
        public
      </span>

      <H6 className="tracking-wider uppercase !text-gray-3 mt-5 flex items-center">
        Direct Messages
        {canSendPrivateMessages && (
          <Button minimal icon={<Icon icon="plus" color={Colors.GRAY3} />} className="ml-2" />
        )}
      </H6>

      <Callout className="mt-2" intent="primary" icon="time">
        Coming soon™
      </Callout>

      {/* {!canSendPrivateMessages && (
        <Callout className="mt-2" intent="warning" icon="warning-sign">
          Set your private message secret to enable direct messages
        </Callout>
      )} */}
    </div>
  )
}

export default BottomMenu
