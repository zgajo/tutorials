class ContactRepository {
  get _baseParams() {
    return {
      TableName: "contacts"
    };
  }

  constructor(documentClient) {
    this._documentClient = documentClient;
  }

  async list() {
    const params = this._createParamObject();
    const response = await this._documentClient.scan(params).promise();

    return response.Items || [];
  }

  async get(id) {
    const params = this._createParamObject({ Key: { id } });
    const response = await this._documentClient.get(params).promise();

    return response.Item;
  }

  async put(contact) {
    const params = this._createParamObject({ Item: contact });
    await this._documentClient.put(params).promise();

    return contact;
  }

  async delete(id) {
    const params = this._createParamObject({ Key: { id } });
    await this._documentClient.delete(params).promise();

    return id;
  }

  _createParamObject(additionalArgs = {}) {
    return Object.assign({}, this._baseParams, additionalArgs);
  }
}

exports.ContactRepository = ContactRepository;
