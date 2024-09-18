import { type Ref, inject } from "vue"

export function useTomoAPI(): TomoAPI {
  const tomoAPI = inject<TomoAPI | null>("tomo-api", null)
  if (tomoAPI === null) {
    throw new Error("useTomoAPI: No outer <tomo-provider /> founded.")
  }
  return tomoAPI
}

export interface TomoAPI {
  openUrl: (url: string) => void
  copyText: (text: string) => void
  version: Ref<string>
}
