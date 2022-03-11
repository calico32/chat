import { Classes, ControlGroup, FormGroup, H1, HTMLSelect, Icon } from '@blueprintjs/core'
import React from 'react'
import { FormState, UseFormRegister } from 'react-hook-form'
import nextConfig from '../../next.config'
import { useWebsocket } from '../../util/context'
import { formErrorMessage } from '../../util/util'
import { UserSettingsFormData } from './UserSettingsDialog'

interface AvatarInputProps {
  errors: FormState<UserSettingsFormData>['errors']
  avatarType: string
  setAvatarType: React.Dispatch<React.SetStateAction<string>>
  register: UseFormRegister<UserSettingsFormData>
}

const AvatarInput: React.VFC<AvatarInputProps> = ({
  avatarType,
  setAvatarType,
  errors,
  register,
}) => {
  const { user } = useWebsocket()

  if (!user) {
    return <H1>Error: not logged in</H1>
  }

  const isUrl = avatarType !== 'gravatar'
  const key = isUrl ? 'avatarUrl' : 'gravatarEmail'
  const error = errors[key]

  return (
    <FormGroup
      label={isUrl ? 'Avatar URL' : 'Avatar'}
      subLabel={isUrl ? <AvatarUrlMessage /> : <GravatarMessage />}
      labelFor="avatar"
      helperText={formErrorMessage(error)}
      intent={error ? 'danger' : 'none'}
    >
      <ControlGroup>
        <HTMLSelect
          onChange={(el) => setAvatarType(el.target.value)}
          value={isUrl ? 'url' : 'gravatar'}
          className="!mr-2"
        >
          <option value="url">URL</option>
          <option value="gravatar">Gravatar</option>
        </HTMLSelect>
        {isUrl ? (
          <input
            key="avatar"
            className={`${Classes.INPUT} ${errors.avatarUrl ? Classes.INTENT_DANGER : ''}`}
            id="avatar"
            placeholder="Avatar URL"
            {...register('avatarUrl', {
              pattern: /^https:\/\/.+/,
            })}
          />
        ) : (
          <input
            key="gravatar"
            className={`${Classes.INPUT} ${errors.gravatarEmail ? Classes.INTENT_DANGER : ''}`}
            id="gravatar"
            placeholder="Gravatar email"
            {...register('gravatarEmail', {
              pattern: {
                value: /^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$/i,
                message: 'Invalid email',
              },
            })}
          />
        )}
      </ControlGroup>
    </FormGroup>
  )
}

const AvatarUrlMessage: React.VFC = () => {
  const domains = nextConfig.images!.domains!.filter((domain) => !domain.startsWith('www.'))

  return (
    <>
      Optional, HTTPS URLs only
      <br />
      Allowed domains: {domains.join(', ')}
    </>
  )
}

const GravatarMessage: React.VFC = () => (
  <>
    <p>Optional, email</p>
    <span className="text-gray-3">
      Set your Gravatar{' '}
      <a
        href="https://gravatar.com/emails/"
        target="_blank"
        rel="noreferrer"
        className="inline-flex items-center"
      >
        here
        <Icon icon="share" size={10} className="ml-1" />
      </a>
      .
    </span>
  </>
)

export default AvatarInput
