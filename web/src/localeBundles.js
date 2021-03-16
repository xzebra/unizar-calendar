/**
 * The bundles here are configured so that each locale only requires loading a single webpack chunk.
 */

const bundles = {
    en: () => import(/* webpackChunkName: "en" */ './locales/en.json'),
    es: () => import(/* webpackChunkName: "es" */ './locales/es.json'),
};

// generate whitelist for i18next
export const availableLocales = Object.keys(bundles);

export default bundles;
