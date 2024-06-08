<template>
  <component :is="currentIcon" v-if="currentIcon" />
</template>

<script>
import { shallowRef } from "vue";

export default {
  props: {
    iconName: {
      type: String,
      required: false,
    },
  },
  setup(props) {
    const currentIcon = shallowRef(null);

    import(`@/components/svg/Svg${props.iconName}.vue`).then(
      (value) => {
        currentIcon.value = value.default;
      },
      (error) => {
        currentIcon.value = "SvgDefault.vue";
      }
    );

    return {
      currentIcon,
    };
  },
};
</script>
