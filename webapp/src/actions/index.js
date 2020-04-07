// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License for license information.
// import {getConfig} from 'mattermost-redux/selectors/entities/general';
// import {PostTypes} from 'mattermost-redux/action_types';
import {id} from '../manifest';
import {PostTypes} from 'mattermost-redux/action_types';

// import Client from '../client';

export function startMeeting(channelId, force = false) {
    // eslint-disable-next-line no-console
    console.log('-----##########################')
    // eslint-disable-next-line no-console
    console.log(String('/plugins/' + id + '/status'))

    return async (dispatch, getState) => {
        console.log('channelid ' + channelId)
        console.log('userid  ' + getState().entities.users.currentUserId)

        fetch('/plugins/' + id + '/status', {
            method: 'GET',
        }).then((response) => {
            return response.json();
        }).then((data) => {
            console.log(data.ocs.data.element);

            // eslint-disable-next-line no-unused-vars
            let rooms = '';
            data.ocs.data.element.forEach((room) => {
                rooms += room.displayName + '\n';
            })

            rooms += '<a class="btn btn-lg btn-primary" href="#">MEETING</a>'

            // const post = {
            //     id: 'nextcloudPlugin' + Date.now(),
            //     create_at: Date.now(),
            //     update_at: 0,
            //     edit_at: 0,
            //     delete_at: 0,
            //     is_pinned: false,
            //     user_id: getState().entities.users.currentUserId,
            //     channel_id: channelId,
            //     root_id: '',
            //     parent_id: '',
            //     original_id: '',
            //     message: rooms,
            //     type: 'system_move_channel',
            //     props: {},
            //     hashtags: '',
            //     pending_post_id: '',
            // };
            //
            // dispatch({
            //     type: PostTypes.RECEIVED_NEW_POST,
            //     data: post,
            //     channelId,
            // });
        });

        // try {
        //
        //     console.log('channelid ' + channelId)
        //     console.log('userid  ' + getState().entities.users.currentUserId)
        //
        //     fetch('/plugins/' + id + '/status', {
        //         method: 'GET',
        //     }).then((response) => {
        //         return response.json();
        //     }).then((data) => {
        //         console.log(data.ocs.data.element);
        //
        //         // eslint-disable-next-line no-unused-vars
        //         let rooms = '';
        //         data.ocs.data.element.forEach((room) => {
        //             rooms += room.displayName + '\n';
        //         })
        //
        //         rooms += '<a class="btn btn-lg btn-primary" href="#">MEETING</a>'
        //
        //         const post = {
        //             id: 'nextcloudPlugin' + Date.now(),
        //             create_at: Date.now(),
        //             update_at: 0,
        //             edit_at: 0,
        //             delete_at: 0,
        //             is_pinned: false,
        //             user_id: getState().entities.users.currentUserId,
        //             channel_id: channelId,
        //             root_id: '',
        //             parent_id: '',
        //             original_id: '',
        //             message: rooms,
        //             type: 'system_move_channel',
        //             props: {},
        //             hashtags: '',
        //             pending_post_id: '',
        //         };
        //
        //         dispatch({
        //             type: PostTypes.RECEIVED_NEW_POST,
        //             data: post,
        //             channelId,
        //         });
        //     });
        // } catch (error) {
        //     let m = 'Ui.... #####################' + error.toString();
        //     if (error.message && error.message[0] === '{') {
        //         const e = JSON.parse(error.message);
        //
        //         // Error is from Zoom API
        //         if (e && e.message) {
        //             m += '\nZoom error: ' + e.message;
        //         }
        //     }
        //
        //     const post = {
        //         id: 'nextcloudPlugin' + Date.now(),
        //         create_at: Date.now(),
        //         update_at: 0,
        //         edit_at: 0,
        //         delete_at: 0,
        //         is_pinned: false,
        //         user_id: getState().entities.users.currentUserId,
        //         channel_id: channelId,
        //         root_id: '',
        //         parent_id: '',
        //         original_id: '',
        //         message: m,
        //         type: 'custom_zoom',
        //         props: {},
        //         hashtags: '',
        //         pending_post_id: '',
        //     };
        //
        //     dispatch({
        //         type: PostTypes.RECEIVED_NEW_POST,
        //         data: post,
        //         channelId,
        //     });
        //
        //     return {error};
        // }

        return {data: true};
    };
}
