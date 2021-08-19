import service from '@/utils/request'

export const createPotential = (data) => {
    return service({
        url: '/potential/createPotential',
        method: 'post',
        data
    })
}

export const deletePotential = (data) => {
    return service({
        url: '/potential/deletePotential',
        method: 'delete',
        data
    })
}

export const updatePotential = (data) => {
    return service({
        url: '/potential/updatePotential',
        method: 'put',
        data
    })
}

export const findPotential = (params) => {
    return service({
        url: '/potential/findPotential',
        method: 'get',
        params
    })
}

export const getPotentialList = (params) => {
    return service({
        url: '/potential/getPotentialList',
        method: 'get',
        params
    })
}

export const upgradePotentialLevel = (data) => {
    return service({
        url: '/potential/upgradePotentialLevel',
        method: 'put',
        data
    })
}

export const degradePotentialLevel = (data) => {
    return service({
        url: '/potential/degradePotentialLevel',
        method: 'put',
        data
    })
}

export const potentialToTrial = (data) => {
    return service({
        url: '/potential/potentialToTrial',
        method: 'put',
        data
    })
}