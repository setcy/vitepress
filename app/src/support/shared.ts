export const EXTERNAL_URL_RE = /^(?:[a-z]+:|\/\/)/i
export const HASH_RE = /#.*$/
export const EXT_RE = /(index)?\.(md|html)$/

export const inBrowser = typeof document !== 'undefined'