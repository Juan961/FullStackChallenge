<template>
  <article class="border rounded px-4 py-2 w-full max-w-md 2xl:max-w-sm overflow-hidden">
    <div class="flex text-left mb-4 items-center">
      <p class="w-3/4 font-medium text-lg">{{ props.result["_source"]["subject"] === '' ? 'No subject' : props.result["_source"]["subject"] }}</p>
      <p class="w-1/4 font-light text-sm">{{ formatDate(props.result["_source"]["date"]) }}</p>
    </div>
    <p>From: <span class="font-light"> {{ props.result["_source"]["from"] }} </span> </p>
    <p>To: <span class="font-light"> {{ props.result["_source"]["to"] }} </span> </p>
    <p>Content: <span class="font-light"> {{ props.result["_source"]["body"] }} </span> </p>
  </article>
</template>

<script setup lang="ts">
interface Email {
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

const props = defineProps<{ result: Email }>();

const formatDate = (dateString: string) => {
  const date = new Date(dateString);

  // Format the date in a user-friendly way
  const formattedDate = date.toLocaleString('en-US', {
    month: 'short',  // Abbreviated month (e.g., Dec)
    day: 'numeric',  // Day of the month (e.g., 30)
    year: 'numeric', // Year (e.g., 2024)
    hour: 'numeric', // Hour (e.g., 6)
    minute: 'numeric', // Minute (e.g., 16)
    hour12: true      // 12-hour clock (AM/PM)
  });

  return formattedDate;
};
</script>
