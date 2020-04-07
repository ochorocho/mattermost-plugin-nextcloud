import React from 'react';

import {id as pluginId} from './manifest';
import {startMeeting} from './actions/index';

import PostTypeZoom from './components/post_type_zoom';
import Icon from './components/Icon.jsx';

export default class Nextcloud {
    initialize(registry, store) {
        registry.registerChannelHeaderButtonAction(
            <Icon/>,
            (channel) => {
                startMeeting(channel.id)(store.dispatch, store.getState);
                console.log(channel);
            },
            'Start Nextcloud Talk Meeting ...'
        );

        registry.registerPostTypeComponent('custom_zoom', PostTypeZoom);
    }

    uninitialize() {
        //eslint-disable-next-line no-console
        console.log(pluginId + '::uninitialize()');
    }
}
