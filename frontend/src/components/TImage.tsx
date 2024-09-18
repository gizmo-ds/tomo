import { PropType, computed, onMounted } from "vue"
import { ImageRenderToolbarProps, NImage, NSkeleton } from "naive-ui"
import { base64ToUint8Array, isVRChatFile } from "@/helper"
import { thumbHashToDataURL } from "thumbhash"
import { FileService } from "@/bindings"

export default defineComponent({
  props: {
    src: String,
    width: [String, Number] as PropType<string | number>,
    height: [String, Number] as PropType<string | number>,
  },
  setup(props) {
    const renderToolbar = ({ nodes }: ImageRenderToolbarProps) => {
      return [
        nodes.rotateCounterclockwise,
        nodes.rotateClockwise,
        nodes.resizeToOriginalSize,
        nodes.zoomOut,
        nodes.zoomIn,
        nodes.close,
      ]
    }

    const isLoading = ref(true)
    const isFailed = ref(false)
    const src = computed(() => {
      if (!isVRChatFile(props.src)) return props.src
      const url = new URL(props.src!)
      return "/@vrchat-file" + url.pathname
    })
    const thumbhash = ref("")
    onMounted(() => {
      if (isVRChatFile(props.src)) {
        FileService.GetThumbHash(src.value!).then((hash) => (thumbhash.value = hash))
      }
    })

    const loading = () => {
      if (thumbhash.value === "") {
        return (
          <NSkeleton width={props.width} height={props.height} style={{ display: "inline-flex" }} />
        )
      } else {
        return (
          <img
            width={props.width}
            height={props.height}
            src={thumbHashToDataURL(base64ToUint8Array(thumbhash.value))}
          />
        )
      }
    }

    return () => [
      <NImage
        src={src.value}
        width={props.width}
        height={props.height}
        renderToolbar={renderToolbar}
        onLoad={() => {
          isLoading.value = false
        }}
        onError={() => {
          isFailed.value = true
          isLoading.value = false
        }}
        style={{
          display: isLoading.value ? "none" : "inline-flex",
        }}
        object-fit="cover"
      />,
      (isLoading.value || isFailed.value) && loading(),
    ]
  },
})
