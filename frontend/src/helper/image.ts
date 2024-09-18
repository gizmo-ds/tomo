export function isVRChatFile(url?: string): boolean {
  if (!url) return false
  return (
    url.startsWith("https://api.vrchat.cloud/api/1/") &&
    /file_[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}/.test(url)
  )
}
