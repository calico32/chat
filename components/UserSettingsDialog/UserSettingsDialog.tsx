import { Button, Classes, ControlGroup, Dialog, FormGroup, H1 } from '@blueprintjs/core'
import React from 'react'
import { SubmitHandler, useForm } from 'react-hook-form'
import { localStorageManager, useWebsocket } from '../../util/context'
import { gravatarUrl, hashSecret } from '../../util/user'
import { formErrorMessage } from '../../util/util'
import AvatarInput from './AvatarInput'
import PrivateMessageSecretInput from './PrivateMessageSecretInput'

interface UserSettingsDialogProps {
  isOpen: boolean
  setOpen: React.Dispatch<React.SetStateAction<boolean>>
}

export interface UserSettingsFormData {
  username: string
  avatarUrl?: string
  gravatarEmail?: string
  secret?: string
}

const UserSettingsDialog: React.VFC<UserSettingsDialogProps> = ({ isOpen, setOpen }) => {
  const { user, updateUser } = useWebsocket()
  const [avatarType, setAvatarType] = React.useState('url')
  const { handleSubmit, register, formState } = useForm<UserSettingsFormData>({
    mode: 'onChange',
    defaultValues: {
      username: user?.username,
      avatarUrl: user?.avatarUrl,
      gravatarEmail: undefined,
      secret: undefined,
    },
  })

  const { errors, isSubmitting } = formState

  if (!user) {
    return (
      <Dialog isOpen={isOpen} onClose={() => setOpen(false)}>
        <H1>Error: not logged in</H1>
      </Dialog>
    )
  }

  const onSubmit: SubmitHandler<UserSettingsFormData> = async (data) => {
    const username = data.username

    let avatarUrl = data.avatarUrl
    if (avatarType === 'gravatar') {
      if (!data.gravatarEmail) {
        console.error('gravatar email required')
        return
      }
      avatarUrl = gravatarUrl(data.gravatarEmail)
    }

    localStorageManager.username = username
    // eslint-disable-next-line @typescript-eslint/prefer-nullish-coalescing
    localStorageManager.avatarUrl = avatarUrl || null

    console.log(data)
    let secretHash: string | undefined
    if (data.secret === 'CLEAR') {
      localStorageManager.secretHash = null
    } else if (data.secret) {
      const hashed = await hashSecret(username, data.secret)
      localStorageManager.secretHash = hashed
      secretHash = hashed
    }

    // await new Promise((resolve) => setTimeout(resolve, 1000))

    updateUser((oldUser) => ({
      ...oldUser!,
      username,
      avatarUrl,
      secretHash: data.secret === 'CLEAR' ? undefined : secretHash ?? oldUser?.secretHash,
    }))

    setOpen(false)
  }

  return (
    <Dialog
      isOpen={isOpen}
      onClose={() => setOpen(false)}
      className="h-full"
      icon="user"
      title="User Settings"
      canEscapeKeyClose={false}
      canOutsideClickClose={false}
    >
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className={Classes.DIALOG_BODY}>
          <FormGroup
            label="Username"
            subLabel={
              <>
                <p>3-32 characters, [A-Za-z0-9_-]</p>
                <span className="text-gray-3">
                  Changing your username will change your identity on the network.
                </span>
              </>
            }
            labelFor="username"
            helperText={formErrorMessage(errors.username)}
            intent={errors.username ? 'danger' : 'none'}
          >
            <input
              className={`${Classes.INPUT} ${errors.username ? Classes.INTENT_DANGER : ''}`}
              id="username"
              placeholder="Username"
              {...register('username', {
                required: true,
                minLength: 3,
                maxLength: 32,
                pattern: /^[A-Za-z0-9_-]+$/,
              })}
            />
          </FormGroup>

          <AvatarInput
            avatarType={avatarType}
            setAvatarType={setAvatarType}
            errors={errors}
            register={register}
          />

          <PrivateMessageSecretInput errors={errors} register={register} />
        </div>
        <div className={Classes.DIALOG_FOOTER}>
          <ControlGroup fill>
            <Button onClick={() => setOpen(false)} disabled={isSubmitting}>
              Cancel
            </Button>
            <Button
              intent="primary"
              type="submit"
              disabled={!formState.isValid || isSubmitting}
              loading={isSubmitting}
            >
              Save
            </Button>
          </ControlGroup>
        </div>
      </form>
    </Dialog>
  )
}

export default UserSettingsDialog
