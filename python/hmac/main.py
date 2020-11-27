import hmac


def main():

    message = "fake_message"
    secrect_key = "fake_key"

    hmac_object = hmac.new(secrect_key.encode(), message.encode(), digestmod="SHA256")
    sig = hmac_object.hexdigest()
    # a5bfb7421f682def4b577d43fe3ef9efab141d192c4fc13df6f325d79e58434d
    print(sig)


if __name__ == "__main__":
    main()
