import { Fragment, provide } from "vue"
import { TomoAPI } from "./TomoAPI"
import { NText } from "naive-ui"
import { Browser, Clipboard } from "@wailsio/runtime"
import { AppService } from "@/bindings"

export default defineComponent({
  setup(_, { slots }) {
    const dialog = useDialog()
    const notification = useNotification()

    const version = ref("develop")
    AppService.GetVersion().then((v) => (version.value = v))

    function copyText(text: string) {
      Clipboard.SetText(text)
        .then(() => notification.success({ content: "Copied to clipboard", duration: 2000 }))
        .catch(() => notification.error({ content: "Failed to copy to clipboard", duration: 2000 }))
    }

    function openUrl(url: string) {
      dialog.success({
        title: "Open External Link",
        content: () => (
          <p>
            Do you want to open the following link?
            <NText code>{url}</NText>
          </p>
        ),
        style: { width: "30rem" },
        showIcon: false,
        positiveText: "Open",
        onPositiveClick: () => Browser.OpenURL(url),
        negativeText: "Copy",
        onNegativeClick: () => copyText(url),
      })
    }

    const api: TomoAPI = { openUrl, copyText, version }
    provide("tomo-api", api)
    return () => <Fragment>{slots.default?.()}</Fragment>
  },
})
