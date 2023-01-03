# TicketCreateReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** |  | 
**Hidden** | Pointer to **float32** |  | [optional] 
**Headers** | Pointer to **string** |  | [optional] 
**ReplyTo** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **float32** |  | [optional] 
**ContactId** | Pointer to **float32** |  | [optional] 
**Cc** | Pointer to **string** |  | [optional] 
**Encrypted** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **float32** |  | [optional] 
**Attachments** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Recipient** | Pointer to **string** |  | [optional] 

## Methods

### NewTicketCreateReply

`func NewTicketCreateReply(body string, ) *TicketCreateReply`

NewTicketCreateReply instantiates a new TicketCreateReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketCreateReplyWithDefaults

`func NewTicketCreateReplyWithDefaults() *TicketCreateReply`

NewTicketCreateReplyWithDefaults instantiates a new TicketCreateReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *TicketCreateReply) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TicketCreateReply) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TicketCreateReply) SetBody(v string)`

SetBody sets Body field to given value.


### GetHidden

`func (o *TicketCreateReply) GetHidden() float32`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *TicketCreateReply) GetHiddenOk() (*float32, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *TicketCreateReply) SetHidden(v float32)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *TicketCreateReply) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetHeaders

`func (o *TicketCreateReply) GetHeaders() string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *TicketCreateReply) GetHeadersOk() (*string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *TicketCreateReply) SetHeaders(v string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *TicketCreateReply) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetReplyTo

`func (o *TicketCreateReply) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *TicketCreateReply) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *TicketCreateReply) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *TicketCreateReply) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSubject

`func (o *TicketCreateReply) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TicketCreateReply) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TicketCreateReply) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TicketCreateReply) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetType

`func (o *TicketCreateReply) GetType() float32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TicketCreateReply) GetTypeOk() (*float32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TicketCreateReply) SetType(v float32)`

SetType sets Type field to given value.

### HasType

`func (o *TicketCreateReply) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContactId

`func (o *TicketCreateReply) GetContactId() float32`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *TicketCreateReply) GetContactIdOk() (*float32, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *TicketCreateReply) SetContactId(v float32)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *TicketCreateReply) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### GetCc

`func (o *TicketCreateReply) GetCc() string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *TicketCreateReply) GetCcOk() (*string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *TicketCreateReply) SetCc(v string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *TicketCreateReply) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetEncrypted

`func (o *TicketCreateReply) GetEncrypted() string`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *TicketCreateReply) GetEncryptedOk() (*string, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *TicketCreateReply) SetEncrypted(v string)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *TicketCreateReply) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetDate

`func (o *TicketCreateReply) GetDate() float32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TicketCreateReply) GetDateOk() (*float32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TicketCreateReply) SetDate(v float32)`

SetDate sets Date field to given value.

### HasDate

`func (o *TicketCreateReply) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetAttachments

`func (o *TicketCreateReply) GetAttachments() []map[string]interface{}`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TicketCreateReply) GetAttachmentsOk() (*[]map[string]interface{}, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TicketCreateReply) SetAttachments(v []map[string]interface{})`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TicketCreateReply) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetRecipient

`func (o *TicketCreateReply) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *TicketCreateReply) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *TicketCreateReply) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *TicketCreateReply) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


