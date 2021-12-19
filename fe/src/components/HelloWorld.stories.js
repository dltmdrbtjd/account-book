import HelloWorld from "@/components/HelloWorld.vue";
export default {
  title: "Example/HelloWorld",
  component: HelloWorld,
};

const Template = (args, { argsTypes }) => ({
  components: { HelloWorld },
  template: `<HelloWorld />`,
});

export const Hello = Template.bind({});
Hello.args = {};
