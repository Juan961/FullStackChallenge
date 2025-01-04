<template>
  <section class="mb-5">
    <h1 class="text-4xl font-medium">Email search</h1>

    <p v-if="error" class="font-light"># records: {{ totalRecords }}</p>
    <p v-if="error" class="font-light">Storage used: {{ storageSize }}MB</p>
  </section>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'

import { info } from '../../actions/info'

const convertBytesToMegabytes = (bytes: number) => bytes / 1024 / 1024;

const totalRecords = ref<null|number>(null)
const storageSize = ref<null|number>(null)
const error = ref<null|string>(null)


onMounted(async () => {
  try {
    const response = await info()

    totalRecords.value = response.doc_num

    const megabytes = convertBytesToMegabytes(response.storage_size)

    storageSize.value = Number(megabytes.toFixed(2))  
  } catch (error) {
    console.error(error)
    error.value = "An error occurred while fetching the information" 
  }
})
</script>
