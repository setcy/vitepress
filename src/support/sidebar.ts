
export interface SidebarLink {
  text: string
  link: string
  docFooterText?: string
}

export type SidebarItem = {
  text: string
  link?: string
  rel?: string
  target?: string
  items?: any[]
  collapsed?: boolean
}