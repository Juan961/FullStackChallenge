<template>
  <main class="container py-8 font-poppins text-[#F6EEE0] min-h-screen">
    <Header />

    <section class="flex mb-5 relative">
      <IconSearch class="w-6 text-[#141313] absolute top-1/2 -translate-y-1/2 left-4" />

      <input class="w-full pl-12 pr-4 py-3 text-[#141313] text-lg border rounded-lg font-light" maxlength="128" type="text" placeholder="Search" v-model="search" @change="newSearch">

      <p v-if="total !== null && took !== null" class="absolute text-[#141313] top-1/2 -translate-y-1/2 right-4 font-light">{{ total }} hits matched in {{ took }}ms</p>
    </section>

    <section v-if="loading" class="flex flex-wrap gap-4 justify-evenly">
      <SearchResultItemSkeleton v-for="i in 6" :key="i" />
    </section>

    <section v-else-if="total !== 0" class="flex flex-wrap gap-4 justify-evenly">
      <SearchResultItem v-for="result in results" :key="result._id" :result="result" />
    </section>

    <section v-else class="text-center">
      <p class="text-xl font-medium">No results found</p>
      <p class="font-light">Sorry if this was not expected</p>
    </section>
  </main>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'

import IconSearch from '../components/icons/Search.vue'

import Header from '../components/search/Header.vue'
import SearchResultItem from '../components/search/ResultItem.vue'
import SearchResultItemSkeleton from '../components/search/ResultItemSkeleton.vue'

import { search as searchAction } from '../actions/search'
import { info } from '../actions/info'

interface Item {
  "@timestamp": string;
  "_id": string;
  "_index": string;
  "_score": number;
  "_source": {
    "@timestamp": string;
    "_id": string;
    "body": string;
    "date": string;
    "from": string;
    "subject": string;
    "to": string;
  };
  "_type": string;
} 

const totalRecords = ref<null|number>(null)
const storageSize = ref<null|number>(null)

const search = ref("")

const took = ref<null|number>(null)
const total = ref<null|number>(null)
const loading = ref(false)
const results = ref<Item[]>([])

const newSearch = async () => {
  loading.value = true

  const response = await searchAction(search.value)

  results.value = response["hits"]
  took.value = response["took"]
  total.value = response["total"]

  loading.value = false
}

const convertBytesToMegabytes = (bytes: number) => bytes / 1024 / 1024;

onMounted(async () => {
  const response = await info()

  totalRecords.value = response.doc_num

  const megabytes = convertBytesToMegabytes(response.storage_size)

  storageSize.value = Number(megabytes.toFixed(2))
})
</script>

<style>
body {
  background-color: #141313;
}
</style>
