<script setup lang="ts">
import TImage from "@/components/TImage"
import { useTomoAPI } from "@/components/tomo"

const tomo = useTomoAPI()
</script>

<template>
  <n-space>
    <n-button @click="tomo.openUrl('https://github.com/gizmo-ds')">{{ tomo.version }}</n-button>
    <div>
      <t-image
        :height="100"
        :width="100"
        src="https://api.vrchat.cloud/api/1/file/file_9164ef8d-c5c7-4fe2-865f-527cfe479d37/3/file"
      />
      <t-image
        :height="100"
        :width="100"
        src="https://api.vrchat.cloud/api/1/file/file_a9cfa444-a5fb-49d1-ac17-0b49e3282e29/1/file"
      />
    </div>
  </n-space>
</template>
