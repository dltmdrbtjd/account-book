import StorybookTest from './StorybookTest.vue';

export default {
  title: 'Example/StorybookTest',
  component: StorybookTest,
};

const Template = () => ({
  component: {StorybookTest},
  template: `<div><StorybookTest>"Hello World"</StorybookTest></div>`,
});

export const Default = Template.bind({});
Default.args = {};
