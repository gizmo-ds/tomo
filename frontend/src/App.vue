<script setup lang="ts">
import App from "./views/App.vue"

import { darkTheme, lightTheme } from "naive-ui"
import { useDark } from "@vueuse/core"
import { computed } from "vue"
import TomoProvider from "./components/tomo"

const isDark = useDark()
const theme = computed(() => (isDark.value ? darkTheme : lightTheme))
const onThemeChange = () => (document.body.style.colorScheme = isDark.value ? "dark" : "light")
watch(isDark, onThemeChange)
onMounted(onThemeChange)

const themeOverrides = computed(
  () =>
    ({
      common: {
        fontWeightStrong: "600",
      },
      Layout: {
        siderColor: isDark.value ? "rgb(16,16,20)" : undefined,
      },
    } as import("naive-ui").GlobalThemeOverrides)
)
</script>

<template>
  <n-config-provider :theme-overrides="themeOverrides" :theme="theme">
    <n-loading-bar-provider>
      <n-notification-provider placement="bottom-right">
        <n-message-provider placement="bottom">
          <n-dialog-provider>
            <tomo-provider>
              <n-layout style="height: 100vh">
                <app />
              </n-layout>
            </tomo-provider>
          </n-dialog-provider>
        </n-message-provider>
      </n-notification-provider>
    </n-loading-bar-provider>
  </n-config-provider>
</template>
