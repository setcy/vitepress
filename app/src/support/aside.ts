import type {Header} from "@/support/shared";

export type MenuItem = Omit<Header, 'slug' | 'children'> & {
    children?: MenuItem[]
}