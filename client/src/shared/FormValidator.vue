<template>
  <div>
    <label v-if="label" :class="labelClass">
      {{ label }} <span v-if="required" class="text-red-600">*</span>
    </label>
    <div class="mt-2">
      <slot></slot>

      <div v-if="errors && errors.length" class="text-red-600">
        <div v-for="(error, index) in errors" :key="index">
          {{ error }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent, computed } from "vue";

import { useGlobalStore } from "@/stores/global";

export default defineComponent({
  props: {
    label: {
      type: String,
      default: null,
    },
    labelClass: {
      type: String,
      default: "block text-sm font-medium leading-6 text-gray-900 form-title",
    },
    name: {
      type: String,
      default: null,
    },
    required: {
      type: Boolean,
      default: false,
    },
  },

  setup(props) {
    const globalStore = useGlobalStore();

    const errors = computed(() => globalStore.getErrors(props.name));

    return {
      errors,
    };
  },
});
</script>

<style lang="scss" scoped>
.form-title {
  font-size: 18px;
}
</style>
