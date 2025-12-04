<script setup lang="ts">
import {error404} from "~/widgets/error/404";
import {error500} from "~/widgets/error/500";

const props = defineProps({
  error: Object
})

const is404 = computed(() => props.error?.statusCode === 404)
</script>

<template>
  <div>
    <error404 v-if="is404"/>
    <error500 v-else/>
  </div>
</template>

<style scoped>

</style>