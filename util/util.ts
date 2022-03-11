import { FieldError } from 'react-hook-form'

export const formErrorMessage = (error: FieldError | undefined): string | undefined => {
  if (!error) return undefined
  if (error.message) return error.message

  switch (error.type) {
    case 'minLength':
      return 'Too short'
    case 'maxLength':
      return 'Too long'
    case 'required':
      return 'Required'
    default:
      return 'Invalid'
  }
}

export const getLocale = (): string => {
  if (navigator.languages !== undefined) {
    return navigator.languages[0]
  }
  return navigator.language
}

export const truncate = (str: string, maxLength: number): string => {
  if (str.length <= maxLength) {
    return str
  }
  return str.substring(0, maxLength) + '...'
}
