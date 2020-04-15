import React from 'react';
import {FormattedMessage} from 'react-intl';

import {MainMenuMobileIcon} from '../../../mattermost-plugin-demo/webapp/src/components/icons';

import {id as pluginId} from './manifest';
import {startMeeting} from './actions/index';

import PostTypeNextcloud from './components/post_type_nextcloud';
import Icon from './components/Icon.jsx';

export default class Nextcloud {
    initialize(registry, store) {
        registry.registerChannelHeaderButtonAction(
            <Icon/>,
            (channel) => {
                startMeeting(channel.id)(store.dispatch, store.getState);
            },
            'Start Nextcloud Talk Meeting ...'
        );

        registry.registerPostTypeComponent('custom_nextcloud', PostTypeNextcloud);

        // registry.registerMainMenuAction(
        //     <FormattedMessage
        //         id='sample.confirmation.dialog'
        //         defaultMessage='Sample Confirmation Dialog'
        //     />,
        //     () => {
        //         window.openInteractiveDialog({
        //             dialog: {
        //                 callback_id: 'nextcloudcallback',
        //                 url: '/plugins/' + pluginId + '/dialog/2',
        //                 title: 'Sample Confirmation Dialog',
        //                 elements: [],
        //                 submit_label: 'Confirm',
        //                 notify_on_cancel: true,
        //                 state: 'somestate',
        //             },
        //         });
        //     },
        //     <MainMenuMobileIcon/>,
        // );
    }

    uninitialize() {
        //eslint-disable-next-line no-console
        console.log(pluginId + '::uninitialize()');
    }
}
