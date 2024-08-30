<script setup lang="ts">
import { AppService } from "../bindings"

const version = ref("")

AppService.GetVersion().then((res) => {
  version.value = res
})
</script>

<template>
  <div>
    <n-button>{{ version }}</n-button>
  </div>
</template>
