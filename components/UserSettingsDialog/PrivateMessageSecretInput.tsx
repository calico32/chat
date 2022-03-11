import { Classes, ControlGroup, FormGroup } from '@blueprintjs/core'
import React from 'react'
import { FormState, UseFormRegister } from 'react-hook-form'
import { formErrorMessage } from '../../util/util'
import { UserSettingsFormData } from './UserSettingsDialog'

interface PrivateMessageSecretProps {
  errors: FormState<UserSettingsFormData>['errors']
  register: UseFormRegister<UserSettingsFormData>
}

const PrivateMessageSecretInput: React.VFC<PrivateMessageSecretProps> = ({ errors, register }) => {
  const [value, setValue] = React.useState('')

  return (
    <FormGroup
      label="Private Message Secret"
      subLabel={
        <>
          <p>Optional, recommended</p>
          <span className="text-gray-3">
            <b>Warning:</b> this is <b>not a password</b> and is only used to verify that you are
            the person who sent/received a message. Do not share this with anyone, and{' '}
            <b>do not use a password</b> you use on other services.
            <br />
            <br />
            Changing your private message secret will change your identity on the network. You will
            lose access to your current private messages.
            <br />
            <br />
            Leave blank to keep your current secret, or enter <code>CLEAR</code> to clear it. If no
            secret is set, you will not be able to send private messages.
          </span>
        </>
      }
      labelFor="secret"
      helperText={
        formErrorMessage(errors.secret) ?? (value === 'CLEAR' && 'Your secret will be cleared')
      }
      intent={errors.username ? 'danger' : 'none'}
    >
      <ControlGroup>
        <input
          className={`${Classes.INPUT} ${errors.secret ? Classes.INTENT_DANGER : ''}`}
          id="secret"
          type={value !== 'CLEAR' ? 'password' : 'text'}
          placeholder="Secret value"
          {...register('secret', {
            onChange: (e: React.ChangeEvent<HTMLInputElement>) => setValue(e.target.value),
            validate: (value) => {
              if (!value) return undefined
              if (value === 'CLEAR') return undefined
              if (value.length < 8) return 'Too short'
              return undefined
            },
          })}
        />
      </ControlGroup>
    </FormGroup>
  )
}

export default PrivateMessageSecretInput
