export * from "./image"

export function base64ToUint8Array(b64Str: string): Uint8Array {
  return Uint8Array.from(globalThis.atob(base64UrlToBase64(b64Str)), (x) => x.charCodeAt(0))
}

function base64UrlToBase64(b64Url: string) {
  return b64Url.replaceAll("-", "+").replaceAll("_", "/")
}
